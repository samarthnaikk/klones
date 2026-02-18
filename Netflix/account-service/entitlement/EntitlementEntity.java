package com.netflix.accountservice.entitlement;

import jakarta.persistence.*;

@Entity
@Table(name = "entitlements")
public class EntitlementEntity {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(nullable = false, unique = true)
    private Long userId;

    @Column(nullable = false)
    private boolean canPlayback = false;

    @Column(nullable = false)
    private boolean canDownload = false;

    @Column(nullable = false)
    private int maxConcurrentStreams = 1;

    public Long getId() { return id; }
    public void setId(Long id) { this.id = id; }

    public Long getUserId() { return userId; }
    public void setUserId(Long userId) { this.userId = userId; }

    public boolean isCanPlayback() { return canPlayback; }
    public void setCanPlayback(boolean canPlayback) { this.canPlayback = canPlayback; }

    public boolean isCanDownload() { return canDownload; }
    public void setCanDownload(boolean canDownload) { this.canDownload = canDownload; }

    public int getMaxConcurrentStreams() { return maxConcurrentStreams; }
    public void setMaxConcurrentStreams(int maxConcurrentStreams) { this.maxConcurrentStreams = maxConcurrentStreams; }
}
