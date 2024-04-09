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
    "suffix" TEXT
);
INSERT INTO "new_User" ("avatarUrl", "coverUrl", "createdAt", "id", "nickname", "password", "slogan", "updatedAt", "username") SELECT "avatarUrl", "coverUrl", "createdAt", "id", "nickname", "password", "slogan", "updatedAt", "username" FROM "User";
DROP TABLE "User";
ALTER TABLE "new_User" RENAME TO "User";
Pragma writable_schema=1;
CREATE UNIQUE INDEX "sqlite_autoindex_users_2" ON "User"("username");
Pragma writable_schema=0;
PRAGMA foreign_key_check;
PRAGMA foreign_keys=ON;
