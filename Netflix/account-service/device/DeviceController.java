package com.netflix.accountservice.device;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/devices")
public class DeviceController {

    private final DeviceService deviceService;

    public DeviceController(DeviceService deviceService) {
        this.deviceService = deviceService;
    }

    // POST /devices/register
    @PostMapping("/register")
    public ResponseEntity<DeviceEntity> registerDevice(@RequestBody DeviceEntity device) {
        return ResponseEntity.ok(deviceService.registerDevice(device));
    }

    // GET /users/{id}/devices
    @GetMapping("/user/{userId}")
    public ResponseEntity<List<DeviceEntity>> getUserDevices(@PathVariable Long userId) {
        return ResponseEntity.ok(deviceService.getUserDevices(userId));
    }

    // DELETE /devices/{id}
    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteDevice(@PathVariable String id) {
        deviceService.deleteDevice(id);
        return ResponseEntity.noContent().build();
    }

    // PATCH /devices/{id}/trust
    @PatchMapping("/{id}/trust")
    public ResponseEntity<DeviceEntity> updateTrust(@PathVariable String id,
                                                    @RequestParam boolean trusted) {
        return ResponseEntity.ok(deviceService.updateTrust(id, trusted));
    }
}
