pub struct EngineCodec;

impl EngineCodec {
     pub fn encode_for_native(data: &str) -> Vec<u8> {
        println!(" Codec: Preparando dados para o núcleo nativo...");
        data.as_bytes().to_vec()
    }
}
