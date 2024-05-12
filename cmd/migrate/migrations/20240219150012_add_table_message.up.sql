-- CreateExtension
CREATE
EXTENSION IF NOT EXISTS "postgis";

-- CreateTable
CREATE TABLE "Message"
(
    "id"         SERIAL       NOT NULL,
    "timestamp"  TIMESTAMP(3) NOT NULL,
    "raster"     TEXT         NOT NULL DEFAULT '',
    "rasterMeta" TEXT         NOT NULL DEFAULT '',
    "polygons"   JSONB        NOT NULL,

    CONSTRAINT "Message_pkey" PRIMARY KEY ("id")
);

CREATE INDEX IF NOT EXISTS idx_polygons
    ON public."Message" USING gin
    (polygons)
    TABLESPACE pg_default;