-- CreateTable
CREATE TABLE `User` (
    `uid` VARCHAR(255) NOT NULL,
    `role` VARCHAR(20) NULL,
    `name` VARCHAR(255) NOT NULL DEFAULT '',
    `createdAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updatedAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `synchedAt` DATETIME(3) NULL,
    `deletedAt` DATETIME(3) NULL,

    PRIMARY KEY (`uid`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `AccessKey` (
    `name` VARCHAR(20) NOT NULL,
    `value` VARCHAR(100) NOT NULL,
    `createdAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updatedAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),

    PRIMARY KEY (`name`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `Import` (
    `id` INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
    `resourceType` VARCHAR(100) NOT NULL,
    `status` VARCHAR(20) NOT NULL,
    `output` TEXT NULL,
    `createdAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updatedAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `objectKey` VARCHAR(255) NOT NULL,

    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `Product` (
    `id` VARCHAR(128) NOT NULL,
    `name` VARCHAR(200) NOT NULL,
    `categoryId1` VARCHAR(30) NOT NULL DEFAULT '',
    `categoryId2` VARCHAR(30) NOT NULL DEFAULT '',
    `categoryId3` VARCHAR(30) NOT NULL DEFAULT '',
    `categoryId4` VARCHAR(30) NOT NULL DEFAULT '',
    `categoryId5` VARCHAR(30) NOT NULL DEFAULT '',
    `note` VARCHAR(1020) NOT NULL DEFAULT '',
    `price` INTEGER UNSIGNED NULL,
    `custom1` VARCHAR(100) NOT NULL DEFAULT '',
    `custom2` VARCHAR(100) NOT NULL DEFAULT '',
    `description` VARCHAR(255) NOT NULL DEFAULT '',
    `details` VARCHAR(1000) NOT NULL DEFAULT '',
    `imageUrl1` VARCHAR(255) NOT NULL DEFAULT '',
    `imageUrl2` VARCHAR(255) NOT NULL DEFAULT '',
    `createdAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updatedAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),

    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `Category` (
    `categoryId1` VARCHAR(50) NOT NULL DEFAULT '',
    `categoryName1` VARCHAR(100) NOT NULL DEFAULT '',
    `categoryId2` VARCHAR(50) NOT NULL DEFAULT '',
    `categoryName2` VARCHAR(100) NOT NULL DEFAULT '',
    `categoryId3` VARCHAR(50) NOT NULL DEFAULT '',
    `categoryName3` VARCHAR(100) NOT NULL DEFAULT '',
    `categoryId4` VARCHAR(50) NOT NULL DEFAULT '',
    `categoryName4` VARCHAR(100) NOT NULL DEFAULT '',
    `categoryId5` VARCHAR(50) NOT NULL DEFAULT '',
    `categoryName5` VARCHAR(100) NOT NULL DEFAULT '',

    UNIQUE INDEX `Category_categoryId1_categoryId2_categoryId3_categoryId4_cat_key`(`categoryId1`, `categoryId2`, `categoryId3`, `categoryId4`, `categoryId5`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `CategoryGroup` (
    `id` VARCHAR(50) NOT NULL,
    `name` VARCHAR(50) NOT NULL,
    `level` INTEGER UNSIGNED NOT NULL,
    `createdAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updatedAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),

    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `CategoryGroupCategory` (
    `categoryId1` VARCHAR(50) NOT NULL DEFAULT '',
    `categoryId2` VARCHAR(50) NOT NULL DEFAULT '',
    `categoryId3` VARCHAR(50) NOT NULL DEFAULT '',
    `categoryId4` VARCHAR(50) NOT NULL DEFAULT '',
    `categoryId5` VARCHAR(50) NOT NULL DEFAULT '',
    `categoryGroupId` VARCHAR(50) NOT NULL,

    UNIQUE INDEX `CategoryGroupCategory_categoryId1_categoryId2_categoryId3_ca_key`(`categoryId1`, `categoryId2`, `categoryId3`, `categoryId4`, `categoryId5`, `categoryGroupId`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `Store` (
    `id` VARCHAR(16) NOT NULL,
    `name` VARCHAR(64) NOT NULL,
    `salesFloorArea` FLOAT NULL,
    `createdAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updatedAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),

    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `UserStore` (
    `email` VARCHAR(255) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `storeId` VARCHAR(16) NOT NULL,
    `createdAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updatedAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),

    UNIQUE INDEX `UserStore_email_storeId_key`(`email`, `storeId`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `Holiday` (
    `date` DATE NOT NULL,
    `name` VARCHAR(255) NOT NULL,

    PRIMARY KEY (`date`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `StoreStatus` (
    `storeId` VARCHAR(16) NOT NULL,
    `date` DATE NOT NULL,
    `isOpen` BOOLEAN NOT NULL DEFAULT false,
    `isExisting` BOOLEAN NOT NULL DEFAULT false,
    `createdAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updatedAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),

    PRIMARY KEY (`storeId`, `date`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- AddForeignKey
ALTER TABLE `CategoryGroupCategory` ADD CONSTRAINT `CategoryGroupCategory_categoryId1_categoryId2_categoryId3_c_fkey` FOREIGN KEY (`categoryId1`, `categoryId2`, `categoryId3`, `categoryId4`, `categoryId5`) REFERENCES `Category`(`categoryId1`, `categoryId2`, `categoryId3`, `categoryId4`, `categoryId5`) ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `CategoryGroupCategory` ADD CONSTRAINT `CategoryGroupCategory_categoryGroupId_fkey` FOREIGN KEY (`categoryGroupId`) REFERENCES `CategoryGroup`(`id`) ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `UserStore` ADD CONSTRAINT `UserStore_storeId_fkey` FOREIGN KEY (`storeId`) REFERENCES `Store`(`id`) ON DELETE RESTRICT ON UPDATE CASCADE;
