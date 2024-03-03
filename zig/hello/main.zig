const std = @import("std");

pub fn main() void {
    // std.debug.print("Hello, {s}!\n", .{"World"});
    const arr = [_]u8{ 'h', 'e', 'l', 'l', 'o' };
    const length = arr.len;
    std.debug.print("{d}", .{length});
}
