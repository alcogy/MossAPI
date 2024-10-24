Deno.serve((req) => {
  const url = new URL(req.url);
  return new Response("Hello: " + url.pathname);
});