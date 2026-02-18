package com.netflix.accountservice.security;

import org.springframework.stereotype.Service;

import java.time.Instant;
import java.util.List;

@Service
public class SecurityService {

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
        double score = events.isEmpty() ? 0.1 : Math.min(events.size() * 0.1, 1.0);
        String level = score < 0.3 ? "LOW" : score < 0.7 ? "MEDIUM" : "HIGH";
        return new SecurityDTO.RiskScoreResponse(userId, score, level);
    }
}
