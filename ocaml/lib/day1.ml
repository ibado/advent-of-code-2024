let num_pair line =
  let l = Util.parse_nums line in
  (List.hd l, List.nth l 1)

let parse_lists input =
  List.map num_pair input
  |> List.fold_left
       (fun (l1, l2) it -> (l1 @ [ fst it ], l2 @ [ snd it ]))
       ([], [])

let part1 input =
  let l1, l2 = parse_lists input in
  let ls1 = List.sort compare l1 in
  let ls2 = List.sort compare l2 in
  List.fold_left2 (fun acc l1 l2 -> acc + abs (l1 - l2)) 0 ls1 ls2

let part2 input =
  let l1, l2 = parse_lists input in
  let freq n l =
    List.fold_left (fun acc it -> if it = n then acc + 1 else acc) 0 l
  in
  List.fold_left (fun acc it -> acc + (it * freq it l2)) 0 l1
