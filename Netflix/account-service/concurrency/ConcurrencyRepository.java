package com.netflix.accountservice.concurrency;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.stereotype.Repository;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Repository
public interface ConcurrencyRepository extends JpaRepository<ConcurrencyEntity, String> {
    List<ConcurrencyEntity> findByUserId(Long userId);

    @Modifying
    @Transactional
    void deleteByStreamId(String streamId);
}
