// types.ts

import { RedisOptions } from 'ioredis';

export type RedisConfig = RedisOptions & {
  port: number;
  host: string;
  password?: string;
};

export type CacheConfig = {
  enabled: boolean;
  redis: RedisConfig;
};

export type CacheRedisConfig = CacheConfig;