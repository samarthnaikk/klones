package com.netflix.accountservice.token;

public class TokenDTO {

    public static class ValidateRequest {
        private String token;

        public String getToken() { return token; }
        public void setToken(String token) { this.token = token; }
    }

    public static class ValidateResponse {
        private boolean valid;
        private Long userId;
        private String email;

        public ValidateResponse(boolean valid, Long userId, String email) {
            this.valid = valid;
            this.userId = userId;
            this.email = email;
        }

        public boolean isValid() { return valid; }
        public Long getUserId() { return userId; }
        public String getEmail() { return email; }
    }

    public static class IntrospectResponse {
        private Long userId;
        private String email;
        private String scope;
        private long expiresIn;

        public IntrospectResponse(Long userId, String email, String scope, long expiresIn) {
            this.userId = userId;
            this.email = email;
            this.scope = scope;
            this.expiresIn = expiresIn;
        }

        public Long getUserId() { return userId; }
        public String getEmail() { return email; }
        public String getScope() { return scope; }
        public long getExpiresIn() { return expiresIn; }
    }
}
