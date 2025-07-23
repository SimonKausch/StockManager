from getbbox import get_bounding_box_from_step
from fastapi import FastAPI, UploadFile, File
import logging
import os

logger = logging.getLogger("uvicorn.error")
logger.setLevel(logging.DEBUG)

app = FastAPI()


@app.get("/test")
async def return_bbox():
    result = get_bounding_box_from_step("model.stp")
    return result


# Return bounding box from stp
#  /box/filepathOfstpFile.stp
#  activate env: conda activate pyoccenv
#  run server: uvicorn main:app --reload
@app.get("/box/{path}")
async def read_item(path: str):
    path = "files/" + path
    result = get_bounding_box_from_step(path)
    return result


@app.post("/upload/")
async def upload_file(file: UploadFile = File(...)):
    # Print information about the file to the console
    logger.debug(f"File Name: {file.filename}")
    logger.debug(f"File MIME Type: {file.content_type}")

    with open("uploaded.stp", "wb") as buffer:
        while True:
            chunk = await file.read(1024 * 8)
            if not chunk:
                break
            buffer.write(chunk)

    # TODO: Get bounding box and delete the file after
    result = get_bounding_box_from_step("uploaded.stp")

    os.remove("uploaded.stp")

    return result
