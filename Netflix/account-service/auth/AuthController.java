package com.netflix.accountservice.auth;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/auth")
public class AuthController {

    private final AuthService authService;

    public AuthController(AuthService authService) {
        this.authService = authService;
    }

    // POST /auth/login
    @PostMapping("/login")
    public ResponseEntity<AuthDTO.TokenResponse> login(@RequestBody AuthDTO.LoginRequest request) {
        return ResponseEntity.ok(authService.login(request.getEmail(), request.getPassword()));
    }

    // POST /auth/logout
    @PostMapping("/logout")
    public ResponseEntity<Void> logout(@RequestParam Long userId) {
        authService.logout(userId);
        return ResponseEntity.noContent().build();
    }

    // POST /auth/refresh
    @PostMapping("/refresh")
    public ResponseEntity<AuthDTO.TokenResponse> refresh(@RequestBody AuthDTO.RefreshRequest request) {
        return ResponseEntity.ok(authService.refresh(request.getRefreshToken()));
    }

    // POST /auth/password-reset/request
    @PostMapping("/password-reset/request")
    public ResponseEntity<Void> requestPasswordReset(@RequestBody AuthDTO.PasswordResetRequest request) {
        authService.requestPasswordReset(request.getEmail());
        return ResponseEntity.accepted().build();
    }

    // POST /auth/password-reset/confirm
    @PostMapping("/password-reset/confirm")
    public ResponseEntity<Void> confirmPasswordReset(@RequestBody AuthDTO.PasswordResetConfirm request) {
        authService.confirmPasswordReset(request.getToken(), request.getNewPassword());
        return ResponseEntity.ok().build();
    }
}
