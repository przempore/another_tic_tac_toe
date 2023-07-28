use tonic::{transport::Server, Request, Response, Status};

use tictactoe::tic_tac_toe_server::{TicTacToe, TicTacToeServer};
use tictactoe::{Board, Cell, Row};

pub mod tictactoe {
    tonic::include_proto!("tictactoe");
}

#[derive(Default)]
struct TicTacToeService {}

#[tonic::async_trait]
impl TicTacToe for TicTacToeService {
    async fn initial_state(&self, _: Request<()>) -> std::result::Result<Response<Board>, Status> {
        let board = Board {
            rows: vec![
                Row {
                    cells: vec![Cell { state: 1 }, Cell { state: 0 }, Cell { state: 0 }],
                },
                Row {
                    cells: vec![Cell { state: 0 }, Cell { state: 1 }, Cell { state: 0 }],
                },
                Row {
                    cells: vec![Cell { state: 0 }, Cell { state: 0 }, Cell { state: 1 }],
                },
            ],
        };

        Ok(Response::new(board))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:50051".parse().unwrap();
    let tic_tac_toe_service = TicTacToeService::default();

    println!("tic_tac_toe listening on {}", addr);

    Server::builder()
        .add_service(TicTacToeServer::new(tic_tac_toe_service))
        .serve(addr)
        .await?;

    Ok(())
}
