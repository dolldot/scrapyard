import os
from dotenv import load_dotenv
from pinecone import Pinecone
from langchain_aws import ChatBedrock
from langchain_aws import BedrockEmbeddings
from langchain_pinecone import PineconeVectorStore
from langchain_community.retrievers import AmazonKnowledgeBasesRetriever

LLAMA3_70B = "meta.llama3-70b-instruct-v1:0"
LLAMA3_8B = "meta.llama3-8b-instruct-v1:0"
LLAMA2_13B = "meta.llama2-13b-chat-v1"


def get_llm():
    profile = os.environ.get("AWS_PROFILE")
    region = os.environ.get("AWS_REGION")

    return ChatBedrock(
        credentials_profile_name=profile,
        region_name=region,
        model_id=LLAMA3_8B,
        model_kwargs=dict(temperature=0),
    )


def get_embeddings():
    profile = os.environ.get("AWS_PROFILE")
    region = os.environ.get("AWS_REGION")

    return BedrockEmbeddings(
        credentials_profile_name=profile,
        region_name=region,
        model_id="amazon.titan-embed-text-v1",
    )


def get_pc_vectorstore():
    api_key = os.environ.get("PINECONE_API_KEY")
    index = os.environ.get("PINECONE_INDEX")

    pc = Pinecone(api_key=api_key)
    index = pc.Index(index)
    embeddings = get_embeddings()
    db = PineconeVectorStore(index, embedding=embeddings)
    return db


def get_aws_knowledge_bases():
    profile = os.environ.get("AWS_PROFILE")
    region = os.environ.get("AWS_REGION")
    aws_knowledge_base_id = os.environ.get("AWS_KNOWLEDGE_BASE_ID")

    return AmazonKnowledgeBasesRetriever(
        credentials_profile_name=profile,
        region_name=region,
        knowledge_base_id=aws_knowledge_base_id,
        retrieval_config={"vectorSearchConfiguration": {"numberOfResults": 1}},
    )
