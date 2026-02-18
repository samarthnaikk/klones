package com.netflix.accountservice.security;

import org.springframework.stereotype.Service;

import java.time.Instant;
import java.util.List;

@Service
public class SecurityService {

    private static final double BASE_RISK_SCORE = 0.1;
    private static final double MAX_RISK_SCORE = 1.0;
    private static final double LOW_RISK_THRESHOLD = 0.3;
    private static final double MEDIUM_RISK_THRESHOLD = 0.7;

    private final SecurityRepository securityRepository;

    public SecurityService(SecurityRepository securityRepository) {
        this.securityRepository = securityRepository;
    }

    public boolean verifyDevice(Long userId, String deviceId, String verificationCode) {
        SecurityEntity event = new SecurityEntity();
        event.setUserId(userId);
        event.setDeviceId(deviceId);
        event.setEventType("DEVICE_VERIFY");
        event.setRiskScore(0.0);
        event.setRecordedAt(Instant.now());
        securityRepository.save(event);
        // In production: validate verificationCode via OTP/email
        return true;
    }

    public SecurityDTO.RiskScoreResponse getRiskScore(Long userId) {
        List<SecurityEntity> events = securityRepository.findByUserId(userId);
        double score = events.isEmpty() ? BASE_RISK_SCORE : Math.min(events.size() * BASE_RISK_SCORE, MAX_RISK_SCORE);
        String level = score < LOW_RISK_THRESHOLD ? "LOW" : score < MEDIUM_RISK_THRESHOLD ? "MEDIUM" : "HIGH";
        return new SecurityDTO.RiskScoreResponse(userId, score, level);
    }
}
