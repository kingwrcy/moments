/*
  Warnings:

  - You are about to alter the column `commentCount` on the `Memo` table. The data in that column could be lost. The data in that column will be cast from `Unsupported("number")` to `Int`.
  - You are about to alter the column `favCount` on the `Memo` table. The data in that column could be lost. The data in that column will be cast from `Unsupported("number")` to `Int`.

*/
-- RedefineTables
PRAGMA foreign_keys=OFF;
CREATE TABLE "new_Memo" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "content" TEXT,
    "imgs" TEXT,
    "favCount" INTEGER NOT NULL DEFAULT 0,
    "commentCount" INTEGER NOT NULL DEFAULT 0,
    "userId" INTEGER NOT NULL,
    "createdAt" DATETIME NOT NULL,
    "updatedAt" DATETIME NOT NULL,
    CONSTRAINT "Memo_userId_fkey" FOREIGN KEY ("userId") REFERENCES "User" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION
);
INSERT INTO "new_Memo" ("commentCount", "content", "createdAt", "favCount", "id", "imgs", "updatedAt", "userId") SELECT coalesce("commentCount", 0) AS "commentCount", "content", "createdAt", coalesce("favCount", 0) AS "favCount", "id", "imgs", "updatedAt", "userId" FROM "Memo";
DROP TABLE "Memo";
ALTER TABLE "new_Memo" RENAME TO "Memo";
PRAGMA foreign_key_check;
PRAGMA foreign_keys=ON;
