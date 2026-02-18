package com.netflix.accountservice.concurrency;

import jakarta.persistence.*;

@Entity
@Table(name = "concurrency_locks")
public class ConcurrencyEntity {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private String lockId;

    @Column(nullable = false)
    private Long userId;

    @Column(nullable = false)
    private String profileId;

    @Column(nullable = false)
    private String streamId;

    public String getLockId() { return lockId; }
    public void setLockId(String lockId) { this.lockId = lockId; }

    public Long getUserId() { return userId; }
    public void setUserId(Long userId) { this.userId = userId; }

    public String getProfileId() { return profileId; }
    public void setProfileId(String profileId) { this.profileId = profileId; }

    public String getStreamId() { return streamId; }
    public void setStreamId(String streamId) { this.streamId = streamId; }
}
