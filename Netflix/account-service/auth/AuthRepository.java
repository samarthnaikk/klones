package com.netflix.accountservice.auth;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.stereotype.Repository;
import org.springframework.transaction.annotation.Transactional;

import java.util.Optional;

@Repository
public interface AuthRepository extends JpaRepository<AuthEntity, Long> {
    Optional<AuthEntity> findByRefreshTokenAndRevokedFalse(String refreshToken);

    @Modifying
    @Transactional
    void deleteByUserId(Long userId);
}
