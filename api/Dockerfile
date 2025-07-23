# Use a base image that includes Miniconda or Anaconda
# Miniconda is often preferred for smaller image sizes
FROM continuumio/miniconda3:latest

# Set the working directory inside the container
WORKDIR /app

# Copy your environment.yml file (if you use one to manage dependencies)
# If you don't have one, you'll install packages directly in the Dockerfile
COPY environment.yml .

# Create the conda environment from environment.yml
# This creates an environment named 'myenv' (or whatever you name it)
RUN conda config --add channels https://conda.anaconda.org/conda-forge
RUN conda env create -f environment.yml

# Activate the environment and install any additional packages if needed
# (though environment.yml is generally preferred for all dependencies)
# RUN conda activate myenv && pip install some-other-package

# Copy your Python script and any other necessary files into the container
COPY main.py .
COPY getbbox.py .
COPY files/ ./files/

EXPOSE 8000

# Set the default command to run when the container starts
# This activates your environment and then runs your Python script
# CMD ["conda", "run", "-n", "pyoccenv", "python", "main.py"]
CMD ["conda", "run", "-n", "pyoccenv", "uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]
# CMD ["conda", "activate", "pyoccenv"]
# CMD ["uvicorn", "main:app", "--reload"]
