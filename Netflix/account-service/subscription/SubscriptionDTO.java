package com.netflix.accountservice.subscription;

import java.time.Instant;

public class SubscriptionDTO {
    private Long id;
    private Long userId;
    private String plan;
    private String status;
    private Instant renewsAt;

    public Long getId() { return id; }
    public void setId(Long id) { this.id = id; }

    public Long getUserId() { return userId; }
    public void setUserId(Long userId) { this.userId = userId; }

    public String getPlan() { return plan; }
    public void setPlan(String plan) { this.plan = plan; }

    public String getStatus() { return status; }
    public void setStatus(String status) { this.status = status; }

    public Instant getRenewsAt() { return renewsAt; }
    public void setRenewsAt(Instant renewsAt) { this.renewsAt = renewsAt; }
}
