const expect = @import("std").testing.expect;

test "if statement" {
    const a = true;
    var x: u16 = 0;
    x += if (a) 1 else 2;

    try expect(x == 1);
}
