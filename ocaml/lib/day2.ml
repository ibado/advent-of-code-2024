let is_asc l =
  let asc n1 n2 =
    let dn = n2 - n1 in
    dn > 0 && dn < 4
  in
  Util.window2 (fun acc n1 n2 -> if acc then Some (asc n1 n2) else None) true l

let is_desc l =
  let desc n1 n2 =
    let dn = n1 - n2 in
    dn > 0 && dn < 4
  in
  Util.window2 (fun acc n1 n2 -> if acc then Some (desc n1 n2) else None) true l

let is_safe l = is_asc l || is_desc l

let is_safe2 l =
  let rec permutations l =
    match l with
    | x :: tl -> [ tl ] @ List.map (fun l -> x :: l) (permutations tl)
    | _ -> []
  in
  List.exists is_safe (permutations l)

let solve f input =
  List.fold_left
    (fun acc it -> if f (Util.parse_nums it) then acc + 1 else acc)
    0 input

let part1 input = solve is_safe input
let part2 input = solve is_safe2 input
