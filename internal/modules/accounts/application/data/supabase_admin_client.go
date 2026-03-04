package data

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	domainErrors "rea/porticos/pkg/errors"
)

type SupabaseAdminClient struct {
	baseURL        string
	serviceRoleKey string
	httpClient     *http.Client
}

func NewSupabaseAdminClient(baseURL, serviceRoleKey string) *SupabaseAdminClient {
	return &SupabaseAdminClient{
		baseURL:        strings.TrimRight(strings.TrimSpace(baseURL), "/"),
		serviceRoleKey: strings.TrimSpace(serviceRoleKey),
		httpClient: &http.Client{
			Timeout: 8 * time.Second,
		},
	}
}

func (c *SupabaseAdminClient) CreateUser(ctx context.Context, email, password string) (string, error) {
	payload := map[string]any{
		"email":         email,
		"password":      password,
		"email_confirm": true,
	}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/auth/v1/admin/users", bytes.NewReader(body))
	if err != nil {
		return "", domainErrors.NewInternalError("SUPABASE_REQUEST_BUILD_ERROR", "no se pudo construir request a Supabase")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", c.serviceRoleKey)
	req.Header.Set("Authorization", "Bearer "+c.serviceRoleKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", domainErrors.NewInternalError("SUPABASE_REQUEST_ERROR", "no se pudo crear usuario en Supabase")
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		if resp.StatusCode == http.StatusConflict || resp.StatusCode == http.StatusUnprocessableEntity || resp.StatusCode == http.StatusBadRequest {
			return "", domainErrors.NewConflictError("ACCOUNT_ALREADY_EXISTS", "email ya registrado en Supabase")
		}
		return "", domainErrors.NewInternalError("SUPABASE_CREATE_USER_ERROR", fmt.Sprintf("error creando usuario en Supabase, status=%d", resp.StatusCode))
	}

	var out struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", domainErrors.NewInternalError("SUPABASE_RESPONSE_ERROR", "respuesta inválida desde Supabase")
	}
	if strings.TrimSpace(out.ID) == "" {
		return "", domainErrors.NewInternalError("SUPABASE_RESPONSE_INVALID", "Supabase no devolvió ID de usuario")
	}

	return out.ID, nil
}

func (c *SupabaseAdminClient) DeleteUser(ctx context.Context, userID string) error {
	userID = strings.TrimSpace(userID)
	if userID == "" {
		return nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, c.baseURL+"/auth/v1/admin/users/"+userID, nil)
	if err != nil {
		return nil
	}
	req.Header.Set("apikey", c.serviceRoleKey)
	req.Header.Set("Authorization", "Bearer "+c.serviceRoleKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	return nil
}
