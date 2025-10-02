use actix::prelude::*;
use actix_web::{get, App, HttpServer, Responder};
use actix_web_actors::ws;
use std::time::Duration;

// 1. `MyWs` must implement `StreamHandler` to handle incoming messages.
// 2. We need to import `StreamHandler` and `Message`.
use actix_web_actors::ws::{Message, ProtocolError};
use actix::StreamHandler;

// A basic WebSocket actor that sends "Hello!" every second.
struct MyWs;

impl Actor for MyWs {
    type Context = ws::WebsocketContext<Self>;

    fn started(&mut self, ctx: &mut Self::Context) {
        ctx.run_interval(Duration::from_secs(1), |_, ctx| {
            ctx.text("Hello!");
        });
    }
}

// 3. We implement the `StreamHandler` trait.
// This is required by `ws::start` to handle messages from the WebSocket stream.
impl StreamHandler<Result<Message, ProtocolError>> for MyWs {
    fn handle(&mut self, msg: Result<Message, ProtocolError>, _ctx: &mut Self::Context) {
        // This is where you would handle incoming messages from the client.
        // For now, we'll just print them.
        match msg {
            Ok(Message::Text(text)) => println!("Received text message: {:?}", text),
            _ => (), // Ignore other message types for now
        }
    }
}

#[get("/ws")]
async fn ws_route(req: actix_web::HttpRequest, stream: actix_web::web::Payload) -> impl Responder {
    ws::start(MyWs, &req, stream)
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| App::new().service(ws_route))
        .bind(("127.0.0.1", 8081))?
        .run()
        .await
}