package utils

import (
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/bsontype"
)

func GetAsInt64(rv bson.RawValue) int64 {
    if rv.Type == bsontype.Int64 {
        return rv.Int64()
    }
    if rv.Type == bsontype.Int32 {
        return int64(rv.Int32())
    }
    return 0
}

func GetAsInt32(rv bson.RawValue) int32 {
    if rv.Type == bsontype.Int32 {
        return rv.Int32()
    }
    if rv.Type == bsontype.Int64 {
        return int32(rv.Int64())
    }
    return 0
}

func GetAsInt(rv bson.RawValue) int {
    if rv.Type == bsontype.Int32 {
        return int(rv.Int32())
    }
    if rv.Type == bsontype.Int64 {
        return int(rv.Int64())
    }
    return 0
}

func GetAsUInt64(rv bson.RawValue) uint64 {
    if rv.Type == bsontype.Int64 {
        return uint64(rv.Int64())
    }
    if rv.Type == bsontype.Int32 {
        return uint64(rv.Int32())
    }
    return 0
}

func GetAsUInt32(rv bson.RawValue) uint32 {
    if rv.Type == bsontype.Int32 {
        return uint32(rv.Int32())
    }
    if rv.Type == bsontype.Int64 {
        return uint32(rv.Int64())
    }
    return 0
}
