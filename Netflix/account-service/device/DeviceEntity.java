package com.netflix.accountservice.device;

import jakarta.persistence.*;

@Entity
@Table(name = "devices")
public class DeviceEntity {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private String id;

    @Column(nullable = false)
    private Long userId;

    @Column(nullable = false)
    private String deviceName;

    @Column(nullable = false)
    private String deviceType;

    @Column(nullable = false)
    private boolean trusted = false;

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
