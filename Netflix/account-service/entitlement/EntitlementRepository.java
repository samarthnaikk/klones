package com.netflix.accountservice.entitlement;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.Optional;

@Repository
public interface EntitlementRepository extends JpaRepository<EntitlementEntity, Long> {
    Optional<EntitlementEntity> findByUserId(Long userId);
}
