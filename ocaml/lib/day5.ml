let parse_input input =
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
  (rules, updates)

let is_correct (n1, n2) l =
  let index_of x l = List.find_index (fun e -> e = x) l in
  match (index_of n1 l, index_of n2 l) with
  | Some in1, Some in2 -> in1 < in2
  | _ -> true

let part1 input =
  let rules, updates = parse_input input in
  updates
  |> List.filter (fun update ->
         List.for_all (fun rule -> is_correct rule update) rules)
  |> List.fold_left (fun acc it -> acc + List.nth it (List.length it / 2)) 0

let part2 input =
  let rules, updates = parse_input input in
  let cmp_pair (a, b) r =
    match List.find_opt (fun (n1, n2) -> n1 = a && n2 = b) r with
    | Some _ -> 1
    | None -> -1
  in
  updates
  |> List.filter (fun update ->
         List.exists (fun rule -> not (is_correct rule update)) rules)
  |> List.map (fun l -> List.sort (fun a b -> cmp_pair (a, b) rules) l)
  |> List.fold_left (fun acc it -> acc + List.nth it (List.length it / 2)) 0
