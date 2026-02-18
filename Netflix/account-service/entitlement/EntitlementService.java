package com.netflix.accountservice.entitlement;

import org.springframework.stereotype.Service;

@Service
public class EntitlementService {

    private final EntitlementRepository entitlementRepository;

    public EntitlementService(EntitlementRepository entitlementRepository) {
        this.entitlementRepository = entitlementRepository;
    }

    public boolean canPlayback(Long userId) {
        return entitlementRepository.findByUserId(userId)
                .map(EntitlementEntity::isCanPlayback)
                .orElse(false);
    }

    public boolean canDownload(Long userId) {
        return entitlementRepository.findByUserId(userId)
                .map(EntitlementEntity::isCanDownload)
                .orElse(false);
    }

    public int getConcurrencyLimit(Long userId) {
        return entitlementRepository.findByUserId(userId)
                .map(EntitlementEntity::getMaxConcurrentStreams)
                .orElse(1);
    }
}
