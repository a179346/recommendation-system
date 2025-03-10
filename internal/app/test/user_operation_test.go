package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/a179346/recommendation-system/internal/app/config"
	"github.com/a179346/recommendation-system/internal/app/database/dbhelper"
	"github.com/a179346/recommendation-system/internal/app/server"
	"github.com/redis/go-redis/v9"
)

func TestUserOperation(t *testing.T) {
	email := "testing@email.com"
	password := "@Ab123"

	db, err := dbhelper.Open()
	if err != nil {
		t.Fatalf("db connection error: %v", err)
	}
	t.Cleanup(func() {
		_, _ = db.Exec(`DELETE FROM user WHERE email = ?`, email)
		db.Close()
	})

	if _, err := db.Exec(`DELETE FROM user WHERE email = ?`, email); err != nil {
		t.Fatalf("db exec error: %v", err)
	}

	redisConfig := config.GetRedisConfig()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
		PoolSize: redisConfig.PoolSize,
	})
	t.Cleanup(func() {
		redisClient.Close()
	})

	server := server.GetServer(db, redisClient)

	t.Run("login should fail before user is registered", func(t *testing.T) {
		body, err := json.Marshal(map[string]any{
			"email":    email,
			"password": password,
		})
		if err != nil {
			t.Errorf("json.Marshal error: %v", err)
			return
		}
		req := httptest.NewRequest("POST", "/api/user/login", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		server.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()
		if got, want := res.StatusCode, 404; got != want {
			respBody, _ := io.ReadAll(res.Body)
			t.Errorf("statusCode: got:%v want:%v %s", got, want, string(respBody))
		}
	})

	t.Run("User register should succeed", func(t *testing.T) {
		body, err := json.Marshal(map[string]any{
			"email":    email,
			"password": password,
		})
		if err != nil {
			t.Errorf("json.Marshal error: %v", err)
			return
		}
		req := httptest.NewRequest("POST", "/api/user/register", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		server.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()
		if got, want := res.StatusCode, 200; got != want {
			respBody, _ := io.ReadAll(res.Body)
			t.Errorf("statusCode: got:%v want:%v %s", got, want, string(respBody))
		}
	})

	t.Run("User register should fail because of duplicate email", func(t *testing.T) {
		body, err := json.Marshal(map[string]any{
			"email":    email,
			"password": password,
		})
		if err != nil {
			t.Errorf("json.Marshal error: %v", err)
			return
		}
		req := httptest.NewRequest("POST", "/api/user/register", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		server.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()
		if got, want := res.StatusCode, 409; got != want {
			respBody, _ := io.ReadAll(res.Body)
			t.Errorf("statusCode: got:%v want:%v %s", got, want, string(respBody))
		}
	})

	t.Run("login should fail before verifying email", func(t *testing.T) {
		body, err := json.Marshal(map[string]any{
			"email":    email,
			"password": password,
		})
		if err != nil {
			t.Errorf("json.Marshal error: %v", err)
			return
		}
		req := httptest.NewRequest("POST", "/api/user/login", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		server.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()
		if got, want := res.StatusCode, 403; got != want {
			respBody, _ := io.ReadAll(res.Body)
			t.Errorf("statusCode: got:%v want:%v %s", got, want, string(respBody))
		}
	})

	t.Run("Verifying email should succeed", func(t *testing.T) {
		row := db.QueryRow("SELECT token FROM user WHERE email = ?", email)
		var i struct{ Token string }
		err := row.Scan(&i.Token)
		if err != nil {
			t.Errorf("query token error: %v", err)
			return
		}

		req := httptest.NewRequest("GET", "/api/user/verify-email?token="+i.Token, nil)
		w := httptest.NewRecorder()

		server.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()
		if got, want := res.StatusCode, 200; got != want {
			respBody, _ := io.ReadAll(res.Body)
			t.Errorf("statusCode: got:%v want:%v %s", got, want, string(respBody))
		}
	})

	t.Run("Verifying email should fail when it's has already been verified", func(t *testing.T) {
		row := db.QueryRow("SELECT token FROM user WHERE email = ?", email)
		var i struct{ Token string }
		err := row.Scan(&i.Token)
		if err != nil {
			t.Errorf("query token error: %v", err)
			return
		}

		req := httptest.NewRequest("GET", "/api/user/verify-email?token="+i.Token, nil)
		w := httptest.NewRecorder()

		server.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()
		if got, want := res.StatusCode, 404; got != want {
			respBody, _ := io.ReadAll(res.Body)
			t.Errorf("statusCode: got:%v want:%v %s", got, want, string(respBody))
		}
	})
	t.Run("login should succeed after verifying email", func(t *testing.T) {
		body, err := json.Marshal(map[string]any{
			"email":    email,
			"password": password,
		})
		if err != nil {
			t.Errorf("json.Marshal error: %v", err)
			return
		}
		req := httptest.NewRequest("POST", "/api/user/login", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		server.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()
		if got, want := res.StatusCode, 200; got != want {
			respBody, _ := io.ReadAll(res.Body)
			t.Errorf("statusCode: got:%v want:%v %s", got, want, string(respBody))
		}
	})
}
