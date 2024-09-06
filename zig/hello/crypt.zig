const std = @import("std");
const crypto = std.crypto;

pub fn main() void {
    // const key = "LH4PkDyQY7pN7nME";
    // const iv = "8RFddpeE3uyx4sFp";
    const message = "hello";

    var outp: u8 = undefined;
    const tst = crypto.hash.sha2.Sha256.hash(message, outp, "iv1");

    std.debug.print("{s}", .{tst});
}
