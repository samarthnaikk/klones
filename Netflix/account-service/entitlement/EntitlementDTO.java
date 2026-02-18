package com.netflix.accountservice.entitlement;

public class EntitlementDTO {
    private Long userId;
    private boolean canPlayback;
    private boolean canDownload;
    private int maxConcurrentStreams;

    public Long getUserId() { return userId; }
    public void setUserId(Long userId) { this.userId = userId; }

    public boolean isCanPlayback() { return canPlayback; }
    public void setCanPlayback(boolean canPlayback) { this.canPlayback = canPlayback; }

    public boolean isCanDownload() { return canDownload; }
    public void setCanDownload(boolean canDownload) { this.canDownload = canDownload; }

    public int getMaxConcurrentStreams() { return maxConcurrentStreams; }
    public void setMaxConcurrentStreams(int maxConcurrentStreams) { this.maxConcurrentStreams = maxConcurrentStreams; }
}
