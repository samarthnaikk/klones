package com.netflix.accountservice.security;

public class SecurityDTO {

    public static class DeviceVerifyRequest {
        private Long userId;
        private String deviceId;
        private String verificationCode;

        public Long getUserId() { return userId; }
        public void setUserId(Long userId) { this.userId = userId; }

        public String getDeviceId() { return deviceId; }
        public void setDeviceId(String deviceId) { this.deviceId = deviceId; }

        public String getVerificationCode() { return verificationCode; }
        public void setVerificationCode(String verificationCode) { this.verificationCode = verificationCode; }
    }

    public static class RiskScoreResponse {
        private Long userId;
        private double score;
        private String level; // LOW, MEDIUM, HIGH

        public RiskScoreResponse(Long userId, double score, String level) {
            this.userId = userId;
            this.score = score;
            this.level = level;
        }

        public Long getUserId() { return userId; }
        public double getScore() { return score; }
        public String getLevel() { return level; }
    }
}
