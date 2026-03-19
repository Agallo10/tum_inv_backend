package storage

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
	"tum_inv_backend/internal/infrastructure/config"
)

// SupabaseStorage gestiona la subida y descarga de archivos en Supabase Storage
type SupabaseStorage struct {
	baseURL    string
	serviceKey string
	bucket     string
	httpClient *http.Client
}

// NewSupabaseStorage crea una nueva instancia del cliente de Supabase Storage
func NewSupabaseStorage(cfg *config.Config) *SupabaseStorage {
	return &SupabaseStorage{
		baseURL:    cfg.SupabaseURL,
		serviceKey: cfg.SupabaseServiceKey,
		bucket:     cfg.SupabaseBucket,
		httpClient: &http.Client{Timeout: 60 * time.Second},
	}
}

// Upload sube un archivo al bucket y retorna la ruta del objeto
func (s *SupabaseStorage) Upload(fileName string, fileData []byte, contentType string) (string, error) {
	url := fmt.Sprintf("%s/storage/v1/object/%s/%s", s.baseURL, s.bucket, fileName)

	req, err := http.NewRequest("POST", url, bytes.NewReader(fileData))
	if err != nil {
		return "", fmt.Errorf("error creando request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+s.serviceKey)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("x-upsert", "true")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error subiendo archivo: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("error de Supabase Storage (status %d): %s", resp.StatusCode, string(body))
	}

	objectPath := fmt.Sprintf("%s/%s", s.bucket, fileName)
	return objectPath, nil
}

// GetPublicURL retorna la URL pública de un archivo
func (s *SupabaseStorage) GetPublicURL(fileName string) string {
	return fmt.Sprintf("%s/storage/v1/object/public/%s/%s", s.baseURL, s.bucket, fileName)
}

// GetSignedURL genera una URL firmada temporal para descargar un archivo privado
func (s *SupabaseStorage) GetSignedURL(fileName string, expiresIn int) (string, error) {
	url := fmt.Sprintf("%s/storage/v1/object/sign/%s/%s", s.baseURL, s.bucket, fileName)

	body := fmt.Sprintf(`{"expiresIn": %d}`, expiresIn)
	req, err := http.NewRequest("POST", url, bytes.NewReader([]byte(body)))
	if err != nil {
		return "", fmt.Errorf("error creando request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+s.serviceKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error obteniendo URL firmada: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("error de Supabase (status %d): %s", resp.StatusCode, string(respBody))
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error leyendo respuesta: %w", err)
	}

	signedPath := extractSignedURL(string(respBody))
	if signedPath == "" {
		return "", fmt.Errorf("no se pudo obtener la URL firmada de la respuesta: %s", string(respBody))
	}

	return s.baseURL + "/storage/v1" + signedPath, nil
}

// extractSignedURL extrae el valor de signedURL del JSON de respuesta
func extractSignedURL(jsonStr string) string {
	key := `"signedURL":"`
	start := bytes.Index([]byte(jsonStr), []byte(key))
	if start == -1 {
		return ""
	}
	start += len(key)
	end := bytes.Index([]byte(jsonStr[start:]), []byte(`"`))
	if end == -1 {
		return ""
	}
	return jsonStr[start : start+end]
}

// Delete elimina un archivo del bucket
func (s *SupabaseStorage) Delete(fileName string) error {
	url := fmt.Sprintf("%s/storage/v1/object/%s/%s", s.baseURL, s.bucket, fileName)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("error creando request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+s.serviceKey)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error eliminando archivo: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error al eliminar (status %d): %s", resp.StatusCode, string(body))
	}

	return nil
}
