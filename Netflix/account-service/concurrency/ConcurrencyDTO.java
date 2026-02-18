package com.netflix.accountservice.concurrency;

public class ConcurrencyDTO {

    public static class CheckRequest {
        private Long userId;
        private String profileId;

        public Long getUserId() { return userId; }
        public void setUserId(Long userId) { this.userId = userId; }

        public String getProfileId() { return profileId; }
        public void setProfileId(String profileId) { this.profileId = profileId; }
    }

    public static class LockRequest {
        private Long userId;
        private String profileId;
        private String streamId;

        public Long getUserId() { return userId; }
        public void setUserId(Long userId) { this.userId = userId; }

        public String getProfileId() { return profileId; }
        public void setProfileId(String profileId) { this.profileId = profileId; }

        public String getStreamId() { return streamId; }
        public void setStreamId(String streamId) { this.streamId = streamId; }
    }
}
