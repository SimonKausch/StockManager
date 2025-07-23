#!/bin/bash
cd "$(dirname "$0")"
source ~/miniconda3/etc/profile.d/conda.sh
conda activate pyoccenv
uvicorn main:app --reload --host 0.0.0.0
