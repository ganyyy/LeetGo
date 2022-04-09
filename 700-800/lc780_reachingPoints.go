package main


func reachingPoints(sx int, sy int, tx int, ty int) bool {

    var max = func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }

    // 逆向思考, 从终点倒推, 只会比当前要小
    for tx > 0 && ty > 0 {

        if tx == sx && ty == sy {
            return true
        }

        // 当二者差距过大时, 可以快速逼近
        if tx > ty {
            tx -= max((tx-sx)/ty, 1) * ty
        } else {
            ty -= max((ty-sy)/tx, 1) * tx
        }
    }


    return false
}