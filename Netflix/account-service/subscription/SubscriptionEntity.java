package com.netflix.accountservice.subscription;

import jakarta.persistence.*;
import java.time.Instant;

@Entity
@Table(name = "subscriptions")
public class SubscriptionEntity {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(nullable = false, unique = true)
    private Long userId;

    @Column(nullable = false)
    private String plan; // BASIC, STANDARD, PREMIUM

    @Column(nullable = false)
    private String status; // ACTIVE, CANCELLED, EXPIRED

    @Column(nullable = false)
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
