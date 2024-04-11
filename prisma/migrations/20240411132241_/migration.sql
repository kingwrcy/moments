-- RedefineTables
PRAGMA foreign_keys=OFF;
CREATE TABLE "new_Memo" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "content" TEXT,
    "imgs" TEXT,
    "favCount" INTEGER NOT NULL DEFAULT 0,
    "commentCount" INTEGER NOT NULL DEFAULT 0,
    "userId" INTEGER NOT NULL,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "music163Url" TEXT,
    "bilibiliUrl" TEXT,
    "location" TEXT,
    "externalUrl" TEXT,
    "externalTitle" TEXT,
    "externalFavicon" TEXT NOT NULL DEFAULT '/favicon.png',
    CONSTRAINT "Memo_userId_fkey" FOREIGN KEY ("userId") REFERENCES "User" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
INSERT INTO "new_Memo" ("bilibiliUrl", "commentCount", "content", "createdAt", "favCount", "id", "imgs", "location", "music163Url", "updatedAt", "userId") SELECT "bilibiliUrl", "commentCount", "content", "createdAt", "favCount", "id", "imgs", "location", "music163Url", "updatedAt", "userId" FROM "Memo";
DROP TABLE "Memo";
ALTER TABLE "new_Memo" RENAME TO "Memo";
PRAGMA foreign_key_check;
PRAGMA foreign_keys=ON;
