use tictactoe::tic_tac_toe_client::TicTacToeClient;

pub mod tictactoe {
    tonic::include_proto!("tictactoe");
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut client = TicTacToeClient::connect("http://[::1]:50051").await?;

    let response = client.initial_state(()).await?;

    println!("RESPONSE: {:?}", response);

    Ok(())
}
