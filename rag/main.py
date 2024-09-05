import os

from dotenv import load_dotenv
from pinecone import Pinecone
from langchain_pinecone import PineconeVectorStore
from langchain_aws import ChatBedrock
from langchain_aws import BedrockEmbeddings
from langchain import hub
from langchain_core.output_parsers import StrOutputParser
from langchain_core.runnables import RunnablePassthrough
from langchain_community.retrievers import AmazonKnowledgeBasesRetriever

load_dotenv()

pinecone_api_key = os.environ.get("PINECONE_API_KEY")
pinecone_index = os.environ.get("PINECONE_INDEX")
aws_profile = os.environ.get("AWS_PROFILE")
aws_region = os.environ.get("AWS_REGION")
aws_knowledge_base_id = os.environ.get("AWS_KNOWLEDGE_BASE_ID")

pc = Pinecone(api_key=pinecone_api_key)

index = pc.Index(pinecone_index)

llama3_70b = "meta.llama3-70b-instruct-v1:0"
llama3_8b = "meta.llama3-8b-instruct-v1:0"
llama2_13b = "meta.llama2-13b-chat-v1"

llm = ChatBedrock(
    credentials_profile_name=aws_profile,
    region_name=aws_region,
    model_id=llama3_8b,
    model_kwargs=dict(temperature=0),
)

embeddings = BedrockEmbeddings(
    credentials_profile_name=aws_profile,
    region_name=aws_region,
    model_id="amazon.titan-embed-text-v1"
)

vector_store = PineconeVectorStore(index=index, embedding=embeddings)
# results = vector_store.similarity_search(query, k=1)
#
# for res in results:
#     print(f"* {res.page_content} [{res.metadata}]")

retriever = AmazonKnowledgeBasesRetriever(
    credentials_profile_name=aws_profile,
    region_name=aws_region,
    knowledge_base_id=aws_knowledge_base_id,
    retrieval_config={
        "vectorSearchConfiguration": {
            "numberOfResults": 1
        }
    },
)

# query = "When is Bella Vista open?"
# results = retriever.invoke(query)

# print(results)


# See full prompt at https://smith.langchain.com/hub/rlm/rag-prompt
prompt = hub.pull("rlm/rag-prompt")


def format_docs(docs):
    return "\n\n".join(doc.page_content for doc in docs)


qa_chain = (
    {
        # "context": vector_store.as_retriever() | format_docs,
        "context": retriever | format_docs,
        "question": RunnablePassthrough(),
    }
    | prompt
    | llm
    | StrOutputParser()
)

result = qa_chain.invoke("Can I bring kids to Bella Vista?")
print(result)
