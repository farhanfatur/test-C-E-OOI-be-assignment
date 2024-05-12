-- CreateTable
CREATE TABLE "User" (
    "id" SERIAL NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,
    "username" TEXT NOT NULL,
    "password" TEXT NOT NULL,

    CONSTRAINT "User_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "MasterPaymentType" (
    "id" SERIAL NOT NULL,
    "name" TEXT NOT NULL,

    CONSTRAINT "MasterPaymentType_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "UserPaymentTypes" (
    "id" SERIAL NOT NULL,
    "userId" INTEGER NOT NULL,
    "masterPaymentTypeId" INTEGER NOT NULL,

    CONSTRAINT "UserPaymentTypes_pkey" PRIMARY KEY ("id")
);

-- AddForeignKey
ALTER TABLE "UserPaymentTypes" ADD CONSTRAINT "UserPaymentTypes_userId_fkey" FOREIGN KEY ("userId") REFERENCES "User"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "UserPaymentTypes" ADD CONSTRAINT "UserPaymentTypes_masterPaymentTypeId_fkey" FOREIGN KEY ("masterPaymentTypeId") REFERENCES "MasterPaymentType"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
