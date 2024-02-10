@DrawImage(kitty, 0, 0)

let x, y = (205, 187)
@ShearAbout(x, y, -0.47, 0.29)

@DrawRoundedRectangle(204, 192, 45, 35, 5)
@ClipPreserve()

let cover = @Resize(cover, 45, 45)
@DrawImage(cover, x, y)

@SetColor(#7a7f5f40)
@FillPreserve()
@SetStrokeWidth(3)
@SetColor(#7a7f5f60)
@Stroke()