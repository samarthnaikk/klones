package com.netflix.accountservice.security;

import jakarta.persistence.*;
import java.time.Instant;

@Entity
@Table(name = "security_events")
public class SecurityEntity {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private String id;

    @Column(nullable = false)
    private Long userId;

    @Column(nullable = false)
    private String eventType; // DEVICE_VERIFY, RISK_SCORE

    @Column(nullable = false)
    private String deviceId;

    @Column(nullable = false)
    private double riskScore;

    @Column(nullable = false)
    private Instant recordedAt;

    public String getId() { return id; }
    public void setId(String id) { this.id = id; }

    public Long getUserId() { return userId; }
    public void setUserId(Long userId) { this.userId = userId; }

    public String getEventType() { return eventType; }
    public void setEventType(String eventType) { this.eventType = eventType; }

    public String getDeviceId() { return deviceId; }
    public void setDeviceId(String deviceId) { this.deviceId = deviceId; }

    public double getRiskScore() { return riskScore; }
    public void setRiskScore(double riskScore) { this.riskScore = riskScore; }

    public Instant getRecordedAt() { return recordedAt; }
    public void setRecordedAt(Instant recordedAt) { this.recordedAt = recordedAt; }
}
