const std = @import("std");
const print = std.debug.print;

// pub fn main() !void {
//     var array_of_string: [4][]const u8 = .{ "1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet" };
//     for (array_of_string) |item| {
//         var word = std.mem.window(u8, item, 1, 1);
//         while (word.next()) |x| {
//             const integer = std.fmt.parseInt(i32, x, 10) catch continue;
//             print("{d}\n", .{integer});
//         }
//     }
// }

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();

    const data = "a1b2c3d4e5f";
    var someval = std.mem.window(u8, data, 1, 1);
    const numb = try allocator.alloc(i32, 10);
    defer allocator.free(numb);
    while (someval.next()) |x| {
        const integer = std.fmt.parseInt(i32, x, 10) catch {
            print("integer: NaN\n", .{});
            continue;
        };

        print("integer: {d}\n", .{integer});
        numb.append(integer);
    }
}
