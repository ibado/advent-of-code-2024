let part1 input =
  let str_not_empty s = s <> "" in
  let rules =
    List.take_while str_not_empty input
    |> List.map (fun s ->
           match String.split_on_char '|' s with
           | [ a; b ] -> (int_of_string a, int_of_string b)
           | _ -> raise (Invalid_argument "unreachable"))
  in
  let updates =
    List.drop_while str_not_empty input
    |> List.drop 1
    |> List.map (fun s -> String.split_on_char ',' s |> List.map int_of_string)
  in
  let is_correct (n1, n2) l =
    let index_of x l = List.find_index (fun e -> e = x) l in
    match (index_of n1 l, index_of n2 l) with
    | Some in1, Some in2 -> in1 < in2
    | _ -> true
  in
  updates
  |> List.filter (fun update ->
         List.for_all (fun rule -> is_correct rule update) rules)
  |> List.fold_left (fun acc it -> acc + List.nth it (List.length it / 2)) 0

let part2 _input = 0
