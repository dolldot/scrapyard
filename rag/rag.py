from langchain import hub
from dotenv import load_dotenv
from langchain_core.runnables import RunnablePassthrough
from langchain_core.output_parsers import StrOutputParser
from utils import get_llm, get_pc_vectorstore


def main():
    load_dotenv()

    query = "Can I bring kids to Bella Vista?"
    rag(query)


def rag(query: str):
    llm = get_llm()
    vectorestore = get_pc_vectorstore()
    # See full prompt at https://smith.langchain.com/hub/rlm/rag-prompt
    prompt = hub.pull("rlm/rag-prompt")
    qa_chain = (
        {
            "context": vectorestore.as_retriever() | format_docs,
            "question": RunnablePassthrough(),
        }
        | prompt
        | llm
        | StrOutputParser()
    )
    response = qa_chain.invoke(query)
    print(response)
    return response


def format_docs(docs):
    return "\n\n".join(doc.page_content for doc in docs)


if __name__ == "__main__":
    main()