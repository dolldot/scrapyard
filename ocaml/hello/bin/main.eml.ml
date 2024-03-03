(* let () = Dream.(run (router [ get "/" (fun (_ : request) -> html Hello.En.v) ])) *)

let render data =
  <html>
  <body>
    <table style="width:100%">
% data |> List.iter(fun (key, value) ->
    <tr>
      <td style="width:40%"><%s key %></td>
      <td style="width:60%"><%s value %></td>
    </tr>
% );
    </table>
  </body>
  </html>

(* let render data = *)
(*   List.iter (fun (key, value) -> *)
(*     print_endline (key ^ ": " ^ value) *)
(*   ) data; *)

let () =
  Dream.run
  @@ Dream.logger
  @@ Dream.router [
    Dream.get "/" (fun _ -> 
      Dream.html (Hello.En.hello "World"));
    Dream.get "/echo" (fun request ->
      let data = Dream.all_headers request in
      Dream.html (render data)
    );
  ]
