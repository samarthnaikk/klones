package com.netflix.accountservice.device;

public class DeviceDTO {
    private String id;
    private Long userId;
    private String deviceName;
    private String deviceType;
    private boolean trusted;

    public String getId() { return id; }
    public void setId(String id) { this.id = id; }

    public Long getUserId() { return userId; }
    public void setUserId(Long userId) { this.userId = userId; }

    public String getDeviceName() { return deviceName; }
    public void setDeviceName(String deviceName) { this.deviceName = deviceName; }

    public String getDeviceType() { return deviceType; }
    public void setDeviceType(String deviceType) { this.deviceType = deviceType; }

    public boolean isTrusted() { return trusted; }
    public void setTrusted(boolean trusted) { this.trusted = trusted; }
}
