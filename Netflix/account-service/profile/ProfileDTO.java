package com.netflix.accountservice.profile;

public class ProfileDTO {
    private Long id;
    private Long userId;
    private String name;
    private String avatarUrl;
    private boolean isKid;

    public Long getId() { return id; }
    public void setId(Long id) { this.id = id; }

    public Long getUserId() { return userId; }
    public void setUserId(Long userId) { this.userId = userId; }

    public String getName() { return name; }
    public void setName(String name) { this.name = name; }

    public String getAvatarUrl() { return avatarUrl; }
    public void setAvatarUrl(String avatarUrl) { this.avatarUrl = avatarUrl; }

    public boolean isKid() { return isKid; }
    public void setKid(boolean kid) { isKid = kid; }
}
