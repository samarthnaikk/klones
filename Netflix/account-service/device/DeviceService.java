package com.netflix.accountservice.device;

import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class DeviceService {

    private final DeviceRepository deviceRepository;

    public DeviceService(DeviceRepository deviceRepository) {
        this.deviceRepository = deviceRepository;
    }

    public DeviceEntity registerDevice(DeviceEntity device) {
        return deviceRepository.save(device);
    }

    public List<DeviceEntity> getUserDevices(Long userId) {
        return deviceRepository.findByUserId(userId);
    }

    public void deleteDevice(String deviceId) {
        deviceRepository.deleteById(deviceId);
    }

    public DeviceEntity updateTrust(String deviceId, boolean trusted) {
        DeviceEntity device = deviceRepository.findById(deviceId)
                .orElseThrow(() -> new RuntimeException("Device not found: " + deviceId));
        device.setTrusted(trusted);
        return deviceRepository.save(device);
    }
}
