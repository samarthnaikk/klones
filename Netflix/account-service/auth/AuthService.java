package com.netflix.accountservice.auth;

import com.netflix.accountservice.user.UserEntity;
import com.netflix.accountservice.user.UserRepository;
import org.springframework.stereotype.Service;

@Service
public class AuthService {

    private final AuthRepository authRepository;
    private final UserRepository userRepository;

    public AuthService(AuthRepository authRepository, UserRepository userRepository) {
        this.authRepository = authRepository;
        this.userRepository = userRepository;
    }

    public AuthDTO.TokenResponse login(String email, String password) {
        UserEntity user = userRepository.findByEmail(email)
                .orElseThrow(() -> new RuntimeException("Invalid credentials"));
        // Password verification would use BCrypt in production
        String accessToken = generateAccessToken(user.getId());
        String refreshToken = generateRefreshToken(user.getId());
        AuthEntity entity = new AuthEntity();
        entity.setUserId(user.getId());
        entity.setRefreshToken(refreshToken);
        authRepository.save(entity);
        return new AuthDTO.TokenResponse(accessToken, refreshToken);
    }

    public void logout(Long userId) {
        authRepository.deleteByUserId(userId);
    }

    public AuthDTO.TokenResponse refresh(String refreshToken) {
        AuthEntity entity = authRepository.findByRefreshTokenAndRevokedFalse(refreshToken)
                .orElseThrow(() -> new RuntimeException("Invalid or expired refresh token"));
        entity.setRevoked(true);
        authRepository.save(entity);
        String newAccess = generateAccessToken(entity.getUserId());
        String newRefresh = generateRefreshToken(entity.getUserId());
        AuthEntity newEntity = new AuthEntity();
        newEntity.setUserId(entity.getUserId());
        newEntity.setRefreshToken(newRefresh);
        authRepository.save(newEntity);
        return new AuthDTO.TokenResponse(newAccess, newRefresh);
    }

    public void requestPasswordReset(String email) {
        // Send password reset email in production
    }

    public void confirmPasswordReset(String token, String newPassword) {
        // Validate token and update password in production
    }

    private String generateAccessToken(Long userId) {
        // JWT generation would be done here using a JwtUtil in production
        return "access-token-for-user-" + userId;
    }

    private String generateRefreshToken(Long userId) {
        return "refresh-token-for-user-" + userId + "-" + System.currentTimeMillis();
    }
}
