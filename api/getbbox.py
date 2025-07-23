import os
import sys

from OCC.Core.Bnd import Bnd_Box

# Corrected import based on your feedback: use lowercase 'brepbndlib'
from OCC.Core.BRepBndLib import brepbndlib
from OCC.Core.IFSelect import IFSelect_RetDone
from OCC.Core.STEPControl import STEPControl_Reader


def get_bounding_box_from_step(step_file_path):
    """
    Reads a STEP file and returns the bounding box dimensions.

    Args:
        step_file_path (str): The path to the STEP file.

    Returns:
        tuple: A tuple containing (min_x, min_y, min_z, max_x, max_y, max_z)
               or None if an error occurred or the file is empty.
    """
    if not os.path.exists(step_file_path):
        print(f"Error: File not found at {step_file_path}", file=sys.stderr)
        return None

    # Initialize STEP reader
    step_reader = STEPControl_Reader()

    # Read the file
    status = step_reader.ReadFile(step_file_path)

    if status != IFSelect_RetDone:
        print(f"Error: Cannot read file {step_file_path}", file=sys.stderr)
        return None

    # Get the number of roots in the file
    nb_roots = step_reader.NbRootsForTransfer()
    if nb_roots <= 0:
        print(f"Warning: No roots found in {step_file_path}", file=sys.stderr)
        return None

    # Transfer the contents of the file
    step_reader.TransferRoots()

    # Get the number of resulting shapes
    nb_shapes = step_reader.NbShapes()
    if nb_shapes <= 0:
        print(
            f"Warning: No shapes found after transfer from {step_file_path}",
            file=sys.stderr,
        )
        return None

    # Get the resulting shape (usually there's one main shape or assembly)
    # You might need to iterate through shapes if the file contains multiple top-level items
    shape = step_reader.Shape(1)  # Assuming the first shape is the one we need

    if shape.IsNull():
        print(f"Error: Transferred shape is null for {step_file_path}", file=sys.stderr)
        return None

    # Compute the bounding box
    bbox = Bnd_Box()
    # Use the correctly cased function from the imported module
    brepbndlib.Add(shape, bbox)  # Using 'brepbndlib' instead of 'BRepBndLib'

    if bbox.IsVoid():
        print(f"Warning: Bounding box is void for {step_file_path}", file=sys.stderr)
        return None

    # Get the min and max coordinates
    min_x, min_y, min_z, max_x, max_y, max_z = bbox.Get()

    return (min_x, min_y, min_z, max_x, max_y, max_z)


# --- How to use the function ---
if __name__ == "__main__":
    # Replace with the actual path to your STP file
    stp_file = "model.stp"

    # Example usage:
    # Create a dummy file for demonstration if it doesn't exist
    # In a real scenario, you would use your actual STP file
    if not os.path.exists(stp_file):
        print(
            f"Creating a dummy file '{stp_file}' for demonstration. Replace this with your actual STP file.",
            file=sys.stderr,
        )
        print(
            "Note: This dummy file is not a valid STEP file and will likely cause the bounding box function to fail.",
            file=sys.stderr,
        )
        with open(stp_file, "w") as f:
            f.write("This is a placeholder. Replace with actual STEP content.\n")

    bbox_coords = get_bounding_box_from_step(stp_file)

    if bbox_coords:
        min_x, min_y, min_z, max_x, max_y, max_z = bbox_coords
        print(f"Bounding Box Coordinates for {stp_file}:")
        print(f"  Min: ({min_x}, {min_y}, {min_z})")
        print(f"  Max: ({max_x}, {max_y}, {max_z})")

        # Calculate dimensions
        length = max_x - min_x
        width = max_y - min_y
        height = max_z - min_z

        print("\nBounding Box Dimensions:")
        print(f"  Length (X): {length}")
        print(f"  Width (Y):  {width}")
        print(f"  Height (Z): {height}")
    else:
        print(f"Could not get bounding box for {stp_file}")
