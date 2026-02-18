package com.netflix.accountservice.concurrency;

import com.netflix.accountservice.entitlement.EntitlementService;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class ConcurrencyService {

    private final ConcurrencyRepository concurrencyRepository;
    private final EntitlementService entitlementService;

    public ConcurrencyService(ConcurrencyRepository concurrencyRepository,
                               EntitlementService entitlementService) {
        this.concurrencyRepository = concurrencyRepository;
        this.entitlementService = entitlementService;
    }

    public boolean check(Long userId, String profileId) {
        int limit = entitlementService.getConcurrencyLimit(userId);
        List<ConcurrencyEntity> locks = concurrencyRepository.findByUserId(userId);
        return locks.size() < limit;
    }

    public ConcurrencyEntity lock(Long userId, String profileId, String streamId) {
        ConcurrencyEntity entity = new ConcurrencyEntity();
        entity.setUserId(userId);
        entity.setProfileId(profileId);
        entity.setStreamId(streamId);
        return concurrencyRepository.save(entity);
    }

    public void release(String streamId) {
        concurrencyRepository.deleteByStreamId(streamId);
    }
}
