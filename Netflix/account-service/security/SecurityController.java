package com.netflix.accountservice.security;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/security")
public class SecurityController {

    private final SecurityService securityService;

    public SecurityController(SecurityService securityService) {
        this.securityService = securityService;
    }

    // POST /security/device-verify
    @PostMapping("/device-verify")
    public ResponseEntity<Boolean> verifyDevice(@RequestBody SecurityDTO.DeviceVerifyRequest request) {
        return ResponseEntity.ok(securityService.verifyDevice(
                request.getUserId(), request.getDeviceId(), request.getVerificationCode()));
    }

    // GET /security/risk-score
    @GetMapping("/risk-score")
    public ResponseEntity<SecurityDTO.RiskScoreResponse> getRiskScore(@RequestParam Long userId) {
        return ResponseEntity.ok(securityService.getRiskScore(userId));
    }
}
