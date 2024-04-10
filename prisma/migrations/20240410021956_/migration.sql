/*
  Warnings:

  - You are about to drop the column `suffix` on the `User` table. All the data in the column will be lost.

*/
-- RedefineTables
PRAGMA foreign_keys=OFF;
CREATE TABLE "new_User" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "username" TEXT NOT NULL,
    "nickname" TEXT,
    "password" TEXT NOT NULL,
    "avatarUrl" TEXT,
    "slogan" TEXT,
    "coverUrl" TEXT,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "enableS3" BOOLEAN NOT NULL DEFAULT false,
    "domain" TEXT,
    "bucket" TEXT,
    "region" TEXT,
    "accessKey" TEXT,
    "secretKey" TEXT,
    "endpoint" TEXT,
    "thumbnailSuffix" TEXT,
    "favicon" TEXT,
    "title" TEXT NOT NULL DEFAULT '极简朋友圈'
);
INSERT INTO "new_User" ("accessKey", "avatarUrl", "bucket", "coverUrl", "createdAt", "domain", "enableS3", "endpoint", "id", "nickname", "password", "region", "secretKey", "slogan", "thumbnailSuffix", "updatedAt", "username") SELECT "accessKey", "avatarUrl", "bucket", "coverUrl", "createdAt", "domain", "enableS3", "endpoint", "id", "nickname", "password", "region", "secretKey", "slogan", "thumbnailSuffix", "updatedAt", "username" FROM "User";
DROP TABLE "User";
ALTER TABLE "new_User" RENAME TO "User";
Pragma writable_schema=1;
CREATE UNIQUE INDEX "sqlite_autoindex_users_2" ON "User"("username");
Pragma writable_schema=0;
PRAGMA foreign_key_check;
PRAGMA foreign_keys=ON;
