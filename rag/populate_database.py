import logging
from uuid import uuid4
from dotenv import load_dotenv
from langchain.text_splitter import RecursiveCharacterTextSplitter
from langchain_community.document_loaders import TextLoader
from langchain.schema.document import Document
from utils import get_pc_vectorstore

# Set up logging
logging.basicConfig(
    level=logging.INFO, format="%(asctime)s - %(levelname)s - %(message)s"
)


def main():
    load_dotenv()

    documents = load_documents()
    chunks = split_documents(documents)
    add_to_pinecone(chunks)


def load_documents():
    loader = TextLoader("./data/bella_vista.txt")
    return loader.load()


def split_documents(documents: list[Document]):
    text_splitter = RecursiveCharacterTextSplitter(
        chunk_size=500, chunk_overlap=0, length_function=len
    )
    return text_splitter.split_documents(documents)


def add_to_pinecone(chunks: list[Document]):
    db = get_pc_vectorstore()
    new_chunks = 0
    existing_chunks = 0

    for chunk in chunks:
        # Check if the document already exists in Pinecone
        existing_docs = db.similarity_search(chunk.page_content, k=1)
        if not existing_docs or existing_docs[0].page_content != chunk.page_content:
            # If the document doesn't exist or is different, add it
            uuid = str(uuid4())
            db.add_documents(documents=[chunk], ids=[uuid])
            new_chunks += 1
            logging.info(f"Added new chunk: {chunk.page_content[:50]}...")
        else:
            existing_chunks += 1
            logging.info(f"Skipped existing chunk: {chunk.page_content[:50]}...")

    logging.info(
        f"Insertion complete. Added {new_chunks} new chunks, skipped {existing_chunks} existing chunks."
    )


if __name__ == "__main__":
    main()
